from django.shortcuts import render, redirect, get_object_or_404
from django.views import View

from .models import Task


class Index(View):
    def get(self, request):
        tasks = Task.objects.all().order_by('-created_at')
        return render(request, 'todo/pages/index.html', {'tasks': tasks})

    def post(self, request):
        title = request.POST.get('task')

        Task.objects.create(title=title)

        return redirect('index')


class Delete(View):
    def get(self, request, pk):
        task = get_object_or_404(Task, pk=pk)
        task.delete()

        tasks = Task.objects.all().order_by('-created_at')
        return render(request, 'todo/pages/index.html', {'tasks': tasks})
