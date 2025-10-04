from django.contrib import admin
from .models import Post, Comment, Tag, Category, Vote


@admin.register(Post)
class PostAdmin(admin.ModelAdmin):
    list_display = ('title', 'author', 'created_at', 'is_published', 'score')
    list_filter = ('is_published', 'category')
    search_fields = ('title', 'body')
    prepopulated_fields = {'slug': ('title',)}


@admin.register(Comment)
class CommentAdmin(admin.ModelAdmin):
    list_display = ('post', 'author', 'created_at', 'is_removed')
    list_filter = ('is_removed',)


admin.site.register(Tag)
admin.site.register(Category)
admin.site.register(Vote)
