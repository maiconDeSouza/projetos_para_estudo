from django.db import models
from django.contrib.auth import get_user_model
from django.db.models import UniqueConstraint

from apps.posts.models import Post
from apps.comments.models import Comment

User = get_user_model()


# Create your models here.
class LikePost(models.Model):
    post = models.ForeignKey(
        Post, on_delete=models.CASCADE, related_name='likes'
    )
    user = models.ForeignKey(User, on_delete=models.CASCADE)

    class Meta:
        constraints = [
            UniqueConstraint(fields=['post', 'user'], name='unique_post_like')
        ]


class LikeComment(models.Model):
    comment = models.ForeignKey(
        Comment, on_delete=models.CASCADE, related_name='likes'
    )
    user = models.ForeignKey(User, on_delete=models.CASCADE)

    class Meta:
        constraints = [
            UniqueConstraint(
                fields=['comment', 'user'], name='unique_comment_like'
            )
        ]
