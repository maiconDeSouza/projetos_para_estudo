from django.urls import path

from . import views

app_name = 'polls'
urlpatterns = [
    path('', views.Index.as_view(), name='index'),
    path('details/<int:pk>/', views.Details.as_view(), name='details'),
]
