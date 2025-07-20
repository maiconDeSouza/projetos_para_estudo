from django.urls import path

from . import views

app_name = 'blog'
urlpatterns = [
    path('', views.PostList.as_view(), name='post_list_view'),
    path(
        'tag/<str:name>',
        views.PostListTag.as_view(),
        name='post_list_tag_view',
    ),
    path(
        'author/<str:name>',
        views.PostListAuthor.as_view(),
        name='post_list_author_view',
    ),
    path('<slug:slug>/', views.PostDetails.as_view(), name='post_datails'),
]
