from django.db import models
from django.utils import timezone

import uuid


# Create your models here.
class Task(models.Model):
    class StatusOptions(models.TextChoices):
        SUCCESS = 'success', 'success'
        PENDING = 'pending', 'pending'
        DANGER = 'danger', 'danger'

    id = models.UUIDField(
        default=uuid.uuid4,
        primary_key=True,
        editable=False,
    )
    title = models.CharField(max_length=100, blank=False, null=False)
    status = models.CharField(
        choices=StatusOptions.choices,
        default=StatusOptions.PENDING,
        max_length=10,
    )
    created_at = models.DateTimeField(default=timezone.now)

    def __str__(self):
        return f'{self.title} - {self.status}'
