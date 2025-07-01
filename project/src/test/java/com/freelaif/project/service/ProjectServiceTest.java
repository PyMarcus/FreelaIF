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

    @Mock
    private SequenceGeneratorService sequenceGeneratorService;

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
    void findByClientById_WhenProjectExists_ReturnsProject(){
        when(projectService.findByClientId(1)).thenReturn(List.of(project));

        List<Project> projectResponse = projectService.findByClientId(1);
        projectResponse.forEach( project1 -> {
                    assertThat(project1.getId()).isEqualTo(1);
                }
        );

    }

    @Test
    void findByDevIdById_WhenProjectExists_ReturnsProject(){
        when(projectService.findByDeveloperId(1)).thenReturn(List.of(project));

        List<Project> projectResponse = projectService.findByDeveloperId(1);
        projectResponse.forEach( project1 -> {
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
        assertThat(_project.getId()).isEqualTo(0);

    }

    @Test
    void update_IfProjectExists_ReturnsProject(){
        when(projectService.associateDeveloper(1, 1)).thenReturn(Optional.of(project));

        Optional<Project> _project1 = projectService.findById(1);
        _project1.ifPresent(
                project1 -> assertThat(project1.getDeveloperId()).isNull()
        );

        Optional<Project> _project = projectService.associateDeveloper(1, 1);

        _project.ifPresent(
                project1 -> assertThat(project1.getDeveloperId()).isEqualTo(1)
        );
    }


}
