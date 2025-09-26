from datetime import timedelta

from django.db.models.signals import post_save
from django.db import transaction
from django.dispatch import receiver

from .models import Project, StudySession


@receiver(post_save, sender=Project)
def create_study_session(sender, instance, created, **kwargs):
    if StudySession.objects.filter(project=instance).exists():
        return

    sessions = []

    for i in range(1, instance.days + 1):
        sessions.append(
            StudySession(
                project=instance,
                user=instance.owner,
                day=i,
                date=instance.start_date + timedelta(days=(i - 1)),
            )
        )

    with transaction.atomic():
        StudySession.objects.bulk_create(sessions)
