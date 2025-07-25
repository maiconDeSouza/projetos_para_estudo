# Generated by Django 5.2.4 on 2025-07-26 01:22

import uuid
from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Todo',
            fields=[
                ('id', models.UUIDField(default=uuid.uuid4, editable=False, primary_key=True, serialize=False)),
                ('task', models.CharField(max_length=100)),
                ('completed', models.BooleanField(default=False)),
            ],
        ),
    ]
