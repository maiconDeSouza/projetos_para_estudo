from django.db import models
from django.contrib.auth import get_user_model
from django.utils.text import slugify

from apps.tags.models import Tag

User = get_user_model()


# Create your models here.
class Post(models.Model):
    author = models.ForeignKey(
        User, on_delete=models.SET_DEFAULT, default='Anônimo'
    )
    title = models.CharField(max_length=255)
    slug = models.SlugField(max_length=300, unique=True)
    cover = models.ImageField(upload_to='cover/%Y/%m/%d/')
    content = models.TextField()  # será um editor de texto
    published = models.BooleanField(default=False)
    tags = models.ForeignKey(Tag, on_delete=models.SET_NULL, null=True)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    def save(self, *args, **kwargs):
        base = slugify(self.title)
        slug = base
        count = 1

        while Post.objects.filter(slug=slug).exists():
            slug = f'{base}-{count}'
            count += 1

        self.slug = slug

        return super().save(*args, **kwargs)

    def __str__(self):
        return f'Autor: {self.author} - {self.title}'
