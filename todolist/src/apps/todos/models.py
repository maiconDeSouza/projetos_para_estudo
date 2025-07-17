from django.db import models


class TaskQuerySet(models.QuerySet):
    def all_task_queries(self):
        tasks_all = self.all()
        count = tasks_all.filter(completed=False).count()

        return {'tasks_all': tasks_all, 'count': count}


class Task(models.Model):
    title = models.CharField(max_length=100)
    completed = models.BooleanField(default=False)

    def __str__(self):
        return self.title

    objects = TaskQuerySet.as_manager()
