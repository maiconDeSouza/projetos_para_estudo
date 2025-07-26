from django.urls import path

from . import views

app_name = 'crud'
urlpatterns = [
    path('add/', views.CRUD.as_view(), name='add'),
    path('del/<uuid:pk>', views.CRUD.as_view(), name='del'),
    path('edit/<uuid:pk>', views.Edit.as_view(), name='edit'),
    path('edition/<uuid:pk>', views.CRUD.as_view(), name='edition'),
    path('completed/<uuid:pk>', views.Completed.as_view(), name='completed'),
]
