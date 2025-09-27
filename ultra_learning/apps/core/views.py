from django.shortcuts import render, redirect
from django.views import View
from django.contrib.auth.mixins import LoginRequiredMixin
from django.contrib.auth.forms import AuthenticationForm
from django.contrib.auth import login
from django.utils import timezone

from .models import Project, StudySession
from .forms import CustomUserCreationForm


class Index(View):
    def get(self, request):
        form = AuthenticationForm()
        context = {
            'form': form,
        }

        return render(request, 'pages/index.html', context)

    def post(self, request):
        form = AuthenticationForm(request, data=request.POST)

        if form.is_valid():
            user = form.get_user()
            login(request, user)
            return redirect('home')

        return redirect('index')


class Home(LoginRequiredMixin, View):
    def get(self, request):
        project = Project.objects.filter(owner=request.user).latest(
            'start_date'
        )

        study_sessions = StudySession.objects.filter(project=project)

        add_up_the_hours_studied = 0

        for study_session in study_sessions:
            add_up_the_hours_studied += study_session.duration_study_session

        total_goal_hours = project.total_goal_minutes // 60
        hours_remaining = (
            project.total_goal_minutes - add_up_the_hours_studied
        ) // 60

        hours_studies = add_up_the_hours_studied // 60

        today = timezone.localdate()

        context = {
            'project': project,
            'total_goal_hours': total_goal_hours,
            'hours_remaining': hours_remaining,
            'hours_studied': hours_studies,
            'study_sessions': study_sessions,
            'today': today,
        }

        return render(request, 'pages/home.html', context)


# projects = Project.objects.filter(owner=request.user)
