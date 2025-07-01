package com.freelaif.project.controller;

import com.freelaif.project.controllers.v1.ProjectController;
import com.freelaif.project.models.Project;
import com.freelaif.project.service.ProjectService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import java.util.Arrays;
import java.util.List;
import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

class ProjectControllerTest {

    @Mock
    private ProjectService projectService;

    @InjectMocks
    private ProjectController projectController;

    private Project sampleProject;

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);

        sampleProject = new Project();
        sampleProject.setId(1);
        sampleProject.setTitle("Test Project");
        sampleProject.setClientId(1);
        sampleProject.setDeveloperId(1);
        sampleProject.setDescription("Sample Description");
        sampleProject.setClientId(123);
    }

    @Test
    void createProject_withValidToken_shouldReturnCreated() {
        when(projectService.save(any(Project.class))).thenReturn(sampleProject);

        ResponseEntity<?> response = projectController.createProject("Bearer valid_token", sampleProject);

        assertEquals(HttpStatus.CREATED, response.getStatusCode());
        assertEquals(sampleProject, response.getBody());
    }

    @Test
    void createProject_withoutToken_shouldReturnUnauthorized() {
        ResponseEntity<?> response = projectController.createProject("", sampleProject);

        assertEquals(HttpStatus.UNAUTHORIZED, response.getStatusCode());
    }

    @Test
    void findAll_withValidToken_shouldReturnProjects() {
        when(projectService.findAll()).thenReturn(Arrays.asList(sampleProject));

        ResponseEntity<?> response = projectController.findAll("Bearer valid_token");

        assertEquals(HttpStatus.OK, response.getStatusCode());
        List<?> body = (List<?>) response.getBody();
        assertNotNull(body);
        assertEquals(1, body.size());
    }

    @Test
    void findById_existingProject_shouldReturnProject() {
        when(projectService.findById(1)).thenReturn(Optional.of(sampleProject));

        ResponseEntity<?> response = projectController.findById(1, "Bearer valid_token");

        assertEquals(HttpStatus.OK, response.getStatusCode());
        assertEquals(sampleProject, response.getBody());
    }

    @Test
    void findById_nonExistingProject_shouldReturnNotFound() {
        when(projectService.findById(99)).thenReturn(Optional.empty());

        ResponseEntity<?> response = projectController.findById(99, "Bearer valid_token");

        assertEquals(HttpStatus.NOT_FOUND, response.getStatusCode());
    }



    @Test
    void findByDeveloperId_ExistingProject_shouldReturnOK() {
        when(projectService.findByDeveloperId(1)).thenReturn(anyList());

        ResponseEntity<?> response = projectController.searchProjects("Test", 1, 1,"token");

        assertEquals(HttpStatus.OK, response.getStatusCode());
    }

    @Test
    void findByTitle_shouldReturnProjects() {
        when(projectService.findProjectByTitle("Test")).thenReturn(Arrays.asList(sampleProject));

        ResponseEntity<?> response = projectController.searchProjects("Test", 1, 1,"token");

        assertEquals(HttpStatus.OK, response.getStatusCode());
        List<?> body = (List<?>) response.getBody();
        assertNotNull(body);
        assertEquals(1, body.size());
    }


}
