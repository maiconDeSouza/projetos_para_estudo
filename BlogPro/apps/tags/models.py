from django.db import models
from django.utils.text import slugify


class Tag(models.Model):
    name = models.CharField(max_length=50, unique=True)
    slug = models.CharField(max_length=60, unique=True, editable=False)

    def _generate_unique_slug(self, base_text):
        base = slugify(base_text) or 'tag'
        slug_candidate = base
        counter = 1
        while (
            Tag.objects.filter(slug=slug_candidate).exclude(pk=self.pk).exists()
        ):
            counter += 1
            slug_candidate = f'{base}-{counter}'

            if len(slug_candidate) > self._meta.get_field('slug').max_length:
                allowed = self._meta.get_field('slug').max_length - len(
                    f'-{counter}'
                )
                slug_candidate = f'{base[:allowed]}-{counter}'
        return slug_candidate

    def save(self, *args, **kwargs):
        if not self.slug:
            self.slug = self._generate_unique_slug(self.name)
        super().save(*args, **kwargs)

    def __str__(self):
        return self.name
