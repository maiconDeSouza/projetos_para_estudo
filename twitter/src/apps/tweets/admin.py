from django.contrib import admin

from apps.tweets.models import Tweet

# Register your models here.


@admin.register(Tweet)
class TweetAdmin(admin.ModelAdmin):
    list_display = ('author', 'content', 'created_at',)
    search_fields = ('author', 'content', 'created_at',)
