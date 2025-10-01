from django.db import models
from django.db.models import Sum
from django.db.models.functions import Coalesce


# Create your models here.
class Question(models.Model):
    question_text = models.CharField(max_length=200)
    pub_date = models.DateTimeField('date published')

    def total_votes(self):
        result = self.choice_set.aggregate(total=Coalesce(Sum('votes'), 0))
        return result['total'] or 0

    @property
    def total_votes_property(self):
        return self.total_votes()

    def __str__(self):
        return self.question_text


class Choice(models.Model):
    question = models.ForeignKey(Question, on_delete=models.CASCADE)
    choice_text = models.CharField(max_length=255)
    votes = models.PositiveIntegerField(default=0)

    def __str__(self):
        return f'{self.question} {self.choice_text}'
