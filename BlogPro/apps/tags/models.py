from django.db import models
from django.utils.text import slugify


# Create your models here.
class Tag(models.Model):
    name = models.CharField(max_length=50, unique=True)
    slug = models.CharField(max_length=60, unique=True, editable=False)

    def save(self, *args, **kwargs):
        self.tag = f'@{slugify(self.name)}'
        super().save(*args, **kwargs)

    def __str__(self):
        return f'{self.tag}'
