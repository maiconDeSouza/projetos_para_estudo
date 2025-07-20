from django.shortcuts import render, get_list_or_404, get_object_or_404
from django.views import View


from .models import Post


# Create your views here.
class PostList(View):
    def get(self, request):
        posts = Post.objects.all().order_by('-published_at')
        return render(request, 'blog/pages/home.html', {'posts': posts})


class PostListTag(View):
    def get(self, request, name):
        posts = get_list_or_404(
            Post.objects.order_by('-published_at'), tags__name=name
        )
        return render(request, 'blog/pages/home.html', {'posts': posts})


class PostListAuthor(View):
    def get(self, request, name):
        posts = get_list_or_404(
            Post.objects.order_by('-published_at'), author__username=name
        )
        return render(request, 'blog/pages/home.html', {'posts': posts})


class PostDetails(View):
    def get(self, request, slug):
        post = get_object_or_404(Post, slug=slug)
        return render(request, 'blog/pages/details.html', {'post': post})
