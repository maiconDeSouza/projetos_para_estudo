from django.db import models
from django.conf import settings
from django.contrib.contenttypes.fields import GenericForeignKey
from django.contrib.contenttypes.models import ContentType
from django.utils.text import slugify
from django.utils import timezone
# Create your models here.


class Tag(models.Model):
    name = models.CharField(max_length=50, unique=True)
    slug = models.SlugField(max_length=50, unique=True)

    class Meta:
        ordering = ['name']

    def __str__(self):
        return self.name


class Category(models.Model):
    name = models.CharField(max_length=90, unique=True)
    slug = models.SlugField(max_length=90, unique=True)

    def __str__(self):
        return self.name


class Post(models.Model):
    author = models.ForeignKey(
        settings.AUTH_USER_MODEL,
        on_delete=models.CASCADE,
        related_name='posts',
    )
    title = models.CharField(max_length=255)
    slug = models.SlugField(max_length=255, unique=True, blank=True)
    body = models.TextField()
    category = models.ForeignKey(
        Category,
        null=True,
        blank=True,
        on_delete=models.SET_NULL,
        related_name='posts',
    )
    tags = models.ManyToManyField(Tag, blank=True, related_name='posts')
    created_at = models.DateTimeField(default=timezone.now, db_index=True)
    updated_at = models.DateTimeField(auto_now=True)
    is_published = models.BooleanField(default=True)
    score = models.PositiveIntegerField(default=0, db_index=True)

    class Meta:
        ordering = ['-score', '-created_at']
        indexes = [
            models.Index(fields=['-created_at']),
            models.Index(fields=['slug']),
        ]

    def save(self, *args, **kwargs):
        if not self.slug:
            base = slugify(self.title)[:240]
            slug = base
            counter = 1

            while Post.objects.filter(slug=slug).exists():
                slug = f'{base}-{counter}'
                counter += 1

            self.slug = slug

        return super().save(*args, **kwargs)

    def __str__(self):
        return self.title


class Comment(models.Model):
    post = models.ForeignKey(
        Post, on_delete=models.CASCADE, related_name='comments'
    )
    author = models.ForeignKey(
        settings.AUTH_USER_MODEL,
        on_delete=models.CASCADE,
        related_name='comments',
    )
    parent = models.ForeignKey(
        'self',
        null=True,
        blank=True,
        on_delete=models.CASCADE,
        related_name='replies',
    )
    body = models.TextField()

    created_at = models.DateTimeField(default=timezone.now, db_index=True)
    updated_at = models.DateTimeField(auto_now=True)
    is_removed = models.BooleanField(default=False)

    class Meta:
        ordering = ['created_at']
        indexes = [models.Index(fields=['post', 'created_at'])]

    def __str__(self):
        return f'Comment by {self.author} on {self.post}'


class Vote(models.Model):
    user = models.ForeignKey(
        settings.AUTH_USER_MODEL, on_delete=models.CASCADE, related_name='votes'
    )
    value = models.SmallIntegerField()
    content_type = models.ForeignKey(ContentType, on_delete=models.CASCADE)
    object_id = models.PositiveIntegerField()
    content_object = GenericForeignKey('content_type', 'object_id')
    created_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        unique_together = (('user', 'content_type', 'object_id'),)
        indexes = [
            models.Index(fields=['content_type', 'object_id']),
            models.Index(fields=['user']),
        ]

    def __str__(self):
        return f'{self.user} -> {self.value} on {self.content_object}'
