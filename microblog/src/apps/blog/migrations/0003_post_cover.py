# Generated by Django 5.2.4 on 2025-07-18 19:59

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('blog', '0002_alter_post_slug'),
    ]

    operations = [
        migrations.AddField(
            model_name='post',
            name='cover',
            field=models.ImageField(blank=True, help_text='Faça upload da imagem de capa aqui.', null=True, upload_to='covers/%Y/%m/%d/', verbose_name='Imagem de Capa'),
        ),
    ]
