from django.shortcuts import render, get_list_or_404, get_object_or_404
from django.views import View


from .models import Post
from .forms import FormsComments


# Create your views here.
class PostList(View):
    def get(self, request):
        posts = Post.objects.all().order_by('-published_at')

        return render(
            request,
            'blog/pages/home.html',
            {'posts': posts},
        )


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
        comments = post.comments.all()
        forms = FormsComments()
        return render(
            request,
            'blog/pages/details.html',
            {'post': post, 'comments': comments, 'forms': forms},
        )

    def post(self, request, slug):
        post = get_object_or_404(Post, slug=slug)
        form = FormsComments(request.POST)
        if form.is_valid() and request.user.is_authenticated:
            comment = form.save(commit=False)
            comment.post = post
            comment.author = request.user
            comment.save()
        comments = post.comments.all()
        forms = FormsComments()
        return render(
            request,
            'blog/pages/details.html',
            {'post': post, 'comments': comments, 'forms': forms},
        )
