from django.shortcuts import render, redirect
from django.contrib.auth.decorators import login_required

from .models import Tweet
from .forms import TweetForm


@login_required
def home(request):
    tweets = Tweet.objects.by_user(request.user)
    context = {'tweets': tweets}
    return render(request, 'tweets/pages/home.html', context)


@login_required
def tweet_create(request):
    if request.method == 'POST':
        form = TweetForm(request.POST)
        if form.is_valid():
            tweet = form.save(commit=False)
            tweet.author = request.user
            tweet.save()
            return redirect('home')
    else:
        form = TweetForm()
    return render(request, 'tweets/pages/tweet_form.html', {'form': form})
