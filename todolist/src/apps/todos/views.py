from django.shortcuts import render, redirect
from django.contrib import messages
from django.views import View

from .models import Task


# Create your views here.
class Home(View):
    def get(self, request):
        data = Task.objects.all_task_queries()
        tasks = data['tasks_all']
        count = data['count']

        return render(
            request, 'todos/pages/home.html', {'tasks': tasks, 'count': count}
        )


class CreateTask(View):
    def post(self, request):
        task = request.POST.get('task')
        if task:
            new_task = Task.objects.create(title=task)
            new_task.save()
            return redirect('home')
        else:
            messages.error(request, 'O preencha uma nova task.')
            return redirect('home')


class ActionTask(View):
    def post(self, request, id):
        task = Task.objects.get(pk=id)
        if 'check' in request.POST:
            task.completed = not task.completed
            task.save()

        if 'trash' in request.POST:
            task.delete()

        return redirect('home')


class ClearCompletedTasks(View):
    def post(self, request):
        task = Task.objects.filter(completed=True)
        task.delete()
        return redirect('home')
