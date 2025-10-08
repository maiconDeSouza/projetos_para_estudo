from django.db import models
from django.contrib.auth import get_user_model

from apps.posts.models import Post
from apps.comments.models import Comment

User = get_user_model()


# Create your models here.
class LikePost(models.Model):
    post = models.ForeignKey(
        Post, on_delete=models.CASCADE, related_name='likes'
    )
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    total = models.PositiveIntegerField(default=0, editable=False)

    def __str__(self):
        return f'{self.post.title} - {self.total}'


class LikeComment(models.Model):
    comment = models.ForeignKey(
        Comment, on_delete=models.CASCADE, related_name='likes'
    )
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    total = models.PositiveIntegerField(default=0, editable=False)

    def __str__(self):
        return f'{self.total}'
