package com.freelaif.project.service;

import com.freelaif.project.models.Project;
import com.freelaif.project.repository.ProjectRepository;
import com.freelaif.project.service.ProjectService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import static org.assertj.core.api.Assertions.assertThat;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.List;
import java.util.Optional;

import static org.mockito.Mockito.when;


@ExtendWith(MockitoExtension.class)
public class ProjectServiceTest {
    private Project project;
    final String title = "Project Test";

    @Mock
    private ProjectRepository projectRepository;

    @InjectMocks
    private ProjectService projectService;

    @BeforeEach
    void setUp(){
        project = new Project();
        project.setId(1);
        project.setTitle(title);
        project.setStatus(Project.Status.PROGRESS);
    }

    @Test
    void findByIdById_WhenProjectExists_ReturnsProject(){
        when(projectService.findById(1)).thenReturn(Optional.of(project));

        Optional<Project> projectResponse = projectService.findById(1);
        projectResponse.ifPresent( project1 -> {
                    assertThat(project1.getId()).isEqualTo(1);
                }
        );

    }

    @Test
    void findByIdByTitle_WhenProjectExists_ReturnsProject(){
        when(projectService.findProjectByTitle(title)).thenReturn(List.of(project));

        List<Project> projects = projectService.findProjectByTitle(title);
        assertThat(projects.size()).isEqualTo(1);

    }

    @Test
    void save_saveProjectExists_ReturnsProject(){
        when(projectService.save(project)).thenReturn(project);

        Project _project = projectService.save(project);
        assertThat(_project).isNotNull();
        assertThat(_project.getId()).isEqualTo(1);

    }
}
