from django.db import models
from django.contrib.auth.models import User
from django.utils.text import slugify

from django_ckeditor_5.fields import CKEditor5Field


class Tag(models.Model):
    name = models.CharField(max_length=30, unique=True)

    def __str__(self):
        return self.name


class Post(models.Model):
    title = models.CharField('Título', max_length=200, unique=True)
    slug = models.SlugField(
        'slug',
        max_length=250,
        unique=True,
        blank=True,
        help_text='Identificador único baseado no título.',
    )
    cover = models.ImageField(
        upload_to='covers/%Y/%m/%d/',
        verbose_name='Imagem de Capa',
        blank=True,
        null=True,
        help_text='Faça upload da imagem de capa aqui.',
    )
    body = CKEditor5Field('Conteúdo', config_name='default')
    author = models.ForeignKey(User, on_delete=models.CASCADE)
    tags = models.ManyToManyField(
        Tag,
        related_name='posts',
        blank=True,
        help_text='Selecione uma ou mais tags.',
    )
    published_at = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return self.title

    def save(self, *args, **kwargs):
        if not self.slug:
            self.slug = slugify(self.title)
        super().save(*args, **kwargs)


class Comment(models.Model):
    post = models.ForeignKey(
        Post, related_name='comments', on_delete=models.CASCADE
    )
    author = models.ForeignKey(User, on_delete=models.CASCADE)
    text = models.TextField('Comentário')
    created_at = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return f'{self.author.username} - {self.post.title}'
