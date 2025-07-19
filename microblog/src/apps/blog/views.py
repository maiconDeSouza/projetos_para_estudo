from django.shortcuts import render
from django.views import View

from .models import Post


# Create your views here.
class PostList(View):
    def get(self, request):
        posts = Post.objects.all()
        return render(request, 'blog/pages/home.html', {'posts': posts})
