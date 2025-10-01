from django.contrib import admin

from .models import Choice, Question
# Register your models here.


class QuestionAdmin(admin.ModelAdmin):
    list_display = ('pk', 'question_text')


admin.site.register(Question, QuestionAdmin)
admin.site.register(Choice)
