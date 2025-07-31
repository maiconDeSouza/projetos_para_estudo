from django.db import models

# Create your models here.


class UserLogin(models.Model):
    name = models.CharField(max_length=255, verbose_name='Nome')
    email = models.EmailField(verbose_name='Email')
    password = models.CharField(max_length=100, verbose_name='Senha')

    def __str__(self):
        return self.name
