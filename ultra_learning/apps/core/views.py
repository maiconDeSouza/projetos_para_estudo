from django.shortcuts import render, redirect
from django.views import View
from django.contrib.auth.mixins import LoginRequiredMixin
from django.contrib.auth.forms import AuthenticationForm
from django.contrib.auth import login

from .models import Project


class Index(View):
    def get(self, request):
        form = AuthenticationForm()
        context = {
            'form': form,
        }

        return render(request, 'pages/index.html', context)


# projects = Project.objects.filter(owner=request.user)
