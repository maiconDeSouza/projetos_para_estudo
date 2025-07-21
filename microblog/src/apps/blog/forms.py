from django import forms

from .models import Comment


class FormsComments(forms.ModelForm):
    class Meta:
        model = Comment
        fields = ['text']
        labels = {
            'text': 'Comentário',
        }
        widgets = {
            'text': forms.Textarea(
                attrs={
                    'rows': 4,
                    'placeholder': 'Escreva seu comentário aqui…',
                    'class': 'input-comment',
                }
            ),
        }
