from django.db import models, IntegrityError, transaction
from django.contrib.auth import get_user_model
from django.utils.text import slugify

from apps.tags.models import Tag

User = get_user_model()


class Post(models.Model):
    author = models.ForeignKey(User, on_delete=models.SET_NULL, null=True)
    title = models.CharField(max_length=255)
    slug = models.SlugField(max_length=300, unique=True, null=True, blank=True)
    cover = models.ImageField(upload_to='cover/%Y/%m/%d/')
    content = models.TextField()
    published = models.BooleanField(default=False)
    tags = models.ManyToManyField(Tag)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    class Meta:
        ordering = ['-created_at']

    def _generate_unique_slug(self, base_text=None):
        base = slugify(base_text or self.title) or 'post'
        max_len = self._meta.get_field('slug').max_length
        slug_candidate = base[:max_len]

        Model = self.__class__
        counter = 1

        while (
            Model.objects.filter(slug=slug_candidate)
            .exclude(pk=self.pk)
            .exists()
        ):
            counter += 1
            suffix = f'-{counter}'
            allowed = max_len - len(suffix)
            slug_candidate = f'{base[:allowed]}{suffix}'

        return slug_candidate

    def save(self, *args, **kwargs):
        if not self.slug:
            self.slug = self._generate_unique_slug()

        tries = 0
        while True:
            try:
                with transaction.atomic():
                    super().save(*args, **kwargs)
                break
            except IntegrityError:
                tries += 1
                if tries > 5:
                    raise

                self.slug = self._generate_unique_slug(f'{self.title}-{tries}')

    def __str__(self):
        return f'Autor: {self.author} - {self.title}'
