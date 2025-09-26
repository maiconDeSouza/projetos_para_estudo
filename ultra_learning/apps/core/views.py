from django.shortcuts import render
from django.views import View
from django.contrib.auth.mixins import LoginRequiredMixin

from .models import Project


class Index(View):
    def get(self, request):
        projects = Project.objects.filter(owner=request.user)

        context = {
            'projects': projects,
        }

        return render(request, 'pages/index.html', context)
