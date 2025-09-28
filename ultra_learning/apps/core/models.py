from datetime import timedelta

from django.db import models
from django.conf import settings


# Create your models here.
class Project(models.Model):
    owner = models.ForeignKey(
        settings.AUTH_USER_MODEL, on_delete=models.CASCADE
    )
    name = models.CharField(max_length=255, verbose_name='Nome do Projeto')
    description = models.TextField(verbose_name='Descrição do Projeto')
    start_date = models.DateField(verbose_name='Data de início do Projeto')
    end_date = models.DateField(
        verbose_name='Data de término do Projeto', null=True, blank=True
    )
    days = models.PositiveIntegerField(
        verbose_name='Quantidade de dias do projeto',
        default=0,
    )
    total_goal_minutes = models.PositiveIntegerField(
        verbose_name='Tempo de estudo em minutos',
        default=0,
    )

    def save(self, *args, **kwargs):
        if self.start_date is None:
            raise ValueError('Data de início do Projeto precisa estar definida')

        self.end_date = self.start_date + timedelta(days=self.days)
        super().save(*args, **kwargs)

    def __str__(self):
        return f'{self.name} ({self.owner})'


class StudySession(models.Model):
    project = models.ForeignKey(Project, on_delete=models.CASCADE)
    user = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.CASCADE)
    day = models.PositiveIntegerField(null=True, blank=True)
    date = models.DateField(null=True, blank=True)
    duration_study_session = models.PositiveIntegerField(
        null=True,
        blank=True,
        verbose_name='Duração do estudo no dia',
        default=0,
    )
    name = models.CharField(max_length=10, null=True, blank=True)
    description = models.CharField(max_length=38, null=True, blank=True)

    def duration_mm_ss(self) -> str:
        minutes = self.duration_study_session // 60
        seconds = self.duration_study_session % 60
        return f'{minutes} Min. {seconds:02d} Seg.'

    def __str__(self):
        return f'Dia {self.day} | {self.date} | {self.project}'
