from django.urls import path

from apps.tweets import views

urlpatterns = [
    path('', views.home, name='home'),
    path('new/', views.tweet_create, name='new'),
]
