from django.shortcuts import render, redirect
from django.views import View
from django.contrib.auth.mixins import LoginRequiredMixin
from django.contrib.auth.forms import AuthenticationForm
from django.contrib.auth import login
from django.utils import timezone

from .models import Project, StudySession
# from .forms import CustomUserCreationForm


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
            add_up_the_hours_studied += (
                study_session.duration_study_session // 60
            )

        total_goal_hours = project.total_goal_minutes // 60

        hours_studies = add_up_the_hours_studied // 60

        percentage_completed = round((hours_studies / total_goal_hours) * 100)

        hours_remaining = total_goal_hours - hours_studies

        today = timezone.localdate()

        context = {
            'project': project,
            'total_goal_hours': total_goal_hours,
            'hours_remaining': hours_remaining,
            'hours_studied': hours_studies,
            'study_sessions': study_sessions,
            'today': today,
            'percentage_completed': percentage_completed,
        }

        return render(request, 'pages/home.html', context)


class UpDurationStudySession(View):
    def post(self, request):
        study_sessions_id = request.POST.get('study-session-id')
        study_session_duration = request.POST.get('study-session-duration')

        study_session = StudySession.objects.get(id=study_sessions_id)
        current = study_session.duration_study_session
        total = current + int(study_session_duration)
        study_session.duration_study_session = total

        study_session.save()

        today = timezone.localdate()

        context = {
            'study_session': study_session,
            'today': today,
        }

        return render(request, 'partials/li.html', context)


class ProgressBar(View):
    def get(self, request):
        project = Project.objects.filter(owner=request.user).latest(
            'start_date'
        )

        study_sessions = StudySession.objects.filter(project=project)

        add_up_the_hours_studied = 0

        for study_session in study_sessions:
            add_up_the_hours_studied += (
                study_session.duration_study_session // 60
            )

        total_goal_hours = project.total_goal_minutes // 60

        hours_studies = add_up_the_hours_studied // 60

        percentage_completed = round((hours_studies / total_goal_hours) * 100)

        context = {
            'percentage_completed': percentage_completed,
        }
        print('Cheguei aqui')
        return render(request, 'partials/progress-bar.html', context)


class MetaProject(View):
    def get(self, request):
        project = Project.objects.filter(owner=request.user).latest(
            'start_date'
        )

        study_sessions = StudySession.objects.filter(project=project)

        add_up_the_hours_studied = 0

        for study_session in study_sessions:
            add_up_the_hours_studied += (
                study_session.duration_study_session // 60
            )

        total_goal_hours = project.total_goal_minutes // 60

        hours_studies = add_up_the_hours_studied // 60

        hours_remaining = total_goal_hours - hours_studies

        context = {
            'total_goal_hours': total_goal_hours,
            'hours_remaining': hours_remaining,
        }

        return render(request, 'partials/meta.html', context)
