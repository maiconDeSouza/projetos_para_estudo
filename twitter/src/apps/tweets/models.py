from django.db import models
from django.contrib.auth.models import User


class TweetQuerySet(models.QuerySet):
    def by_user(self, user):
        return self.filter(author=user)


class Tweet(models.Model):
    author = models.ForeignKey(User, on_delete=models.CASCADE)
    content = models.CharField(max_length=280)
    created_at = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return f"{self.author.username}: {self.content[:20]}"

    objects = TweetQuerySet.as_manager()
