package com.freelaif.project.repository;

import com.freelaif.project.models.Project;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import static org.assertj.core.api.Assertions.assertThat;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.List;

import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
public class ProjectRepositoryTest {
    @Mock
    private ProjectRepository projectRepository;
    private Project project;
    final private String title = "Projeto Teste";


    @BeforeEach
    void setUp() {
        project = new Project();
        project.setId(1);
        project.setTitle(title);
        project.setStatus(Project.Status.PROGRESS);
    }

    @Test
    void getProjectById_WhenProjectExists_ReturnsProject() {
        when(projectRepository.save(this.project)).thenReturn(this.project);

        Project savedProject = projectRepository.save(project);

        assertThat(savedProject).isNotNull();
        assertThat(savedProject.getId()).isEqualTo(1);
    }

    @Test
    void findByTitle_WhenProjectExists_ReturnsProjectList(){
        when(projectRepository.findByTitle(title)).thenReturn(List.of(project));

        List<Project> projects = projectRepository.findByTitle(title);

        assertThat(projects).isNotEmpty();
        assertThat(projects.get(0).getId()).isEqualTo(1);
    }
}
