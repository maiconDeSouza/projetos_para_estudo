from django.urls import path

from . import views

urlpatterns = [
    path('', views.Home.as_view(), name='home'),
    path('create_task/', views.CreateTask.as_view(), name='create_task'),
    path('action/<int:id>', views.ActionTask.as_view(), name='action'),
]
