from django.contrib import admin

from .models import Project, StudySession


# Register your models here.
class ProjectAdmin(admin.ModelAdmin):
    list_display = ('name', 'days', 'start_date', 'end_date')
    search_fields = ('name',)


class StudySessionAdmin(admin.ModelAdmin):
    list_display = ('project', 'user', 'name')


admin.site.register(Project, ProjectAdmin)
admin.site.register(StudySession, StudySessionAdmin)
