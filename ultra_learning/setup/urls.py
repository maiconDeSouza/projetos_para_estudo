from django.contrib import admin
from django.urls import path

from apps.core import views

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', views.Index.as_view(), name='index'),
    path('home/', views.Home.as_view(), name='home'),
    path('duration/', views.UpDurationStudySession.as_view(), name='duration'),
    path('progress/', views.ProgressBar.as_view(), name='progress'),
    path('meta/', views.MetaProject.as_view(), name='meta'),
]
