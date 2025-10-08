from django.db import models
from django.contrib.auth import get_user_model

from apps.posts.models import Post

User = get_user_model()


# Create your models here.
class Comment(models.Model):
    author = models.ForeignKey(User, on_delete=models.SET_NULL, null=True)
    post = models.ForeignKey(
        Post, on_delete=models.CASCADE, related_name='posts'
    )
    content = models.TextField()
    created_at = models.DateTimeField(auto_now_add=True)
