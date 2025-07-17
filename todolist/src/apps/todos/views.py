from django.shortcuts import render, redirect
from django.views import View

from .models import Taks


# Create your views here.
class Home(View):
    def get(self, request):
        tasks = Taks.objects.all()
        count = Taks.objects.filter(completed=False).count()
        return render(
            request, 'todos/pages/home.html', {'tasks': tasks, 'count': count}
        )


class CreateTask(View):
    def post(self, request):
        task = request.POST.get('task')
        new_task = Taks.objects.create(title=task)
        new_task.save()
        return redirect('home')


class ActionTask(View):
    def post(self, request, id):
        task = Taks.objects.get(pk=id)
        if 'check' in request.POST:
            task.completed = not task.completed
            task.save()

        if 'trash' in request.POST:
            task.delete()

        return redirect('home')
