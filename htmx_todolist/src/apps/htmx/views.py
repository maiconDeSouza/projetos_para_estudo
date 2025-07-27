from urllib.parse import parse_qs

from django.shortcuts import render, HttpResponse
from django.shortcuts import get_object_or_404
from django.views import View

from apps.todo.models import Todo


class CreateTask(View):
    def post(self, request):
        task = request.POST.get('task')

        if task:
            todo = Todo.objects.create(task=task)
            return render(
                request, 'todo/partials/list-task.html', {'todo': todo}
            )


class DeleteTask(View):
    def delete(self, request, pk):
        todo = get_object_or_404(Todo, pk=pk)
        todo.delete()
        return HttpResponse('')


class Edit(View):
    def get(self, request, pk):
        todo = get_object_or_404(Todo, pk=pk)
        return render(request, 'todo/partials/edit-task.html', {'todo': todo})


class EditionTask(View):
    def put(self, request, pk):
        todo = get_object_or_404(Todo, pk=pk)
        data = parse_qs(request.body.decode())
        edition_task = data.get('edition-task', [''])[0]
        todo.task = edition_task
        todo.save()
        return render(request, 'todo/partials/list-task.html', {'todo': todo})


class Completed(View):
    def get(self, request, pk):
        todo = get_object_or_404(Todo, pk=pk)
        todo.completed = not todo.completed
        todo.save()
        return render(request, 'todo/partials/list-task.html', {'todo': todo})


class FilterTasks(View):
    def get(self, request, filter_type):
        todos = Todo.objects.all()
        if filter_type == 'completed':
            todos = Todo.objects.filter(completed=True)
        elif filter_type == 'pending':
            todos = Todo.objects.filter(completed=False)

        return render(
            request, 'todo/partials/list-task-list.html', {'todos': todos}
        )
