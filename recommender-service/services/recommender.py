import random


def generate_recommendations():
    # Simulate recommendations
    projects = [
        {'title': 'Projeto Web', 'description': 'HTML, CSS, JavaScript', 'link': '#'},
        {'title': 'App Mobile', 'description': 'React Native, Firebase', 'link': '#'},
        {'title': 'Design UX', 'description': 'Figma, Pesquisa UX', 'link': '#'}
    ]
    return random.sample(projects, 2)
