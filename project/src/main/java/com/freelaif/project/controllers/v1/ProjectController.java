package com.freelaif.project.controllers.v1;

import com.freelaif.project.models.Project;
import com.freelaif.project.service.ProjectService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/api/v1/projects")
public class ProjectController {
    @Autowired
    ProjectService projectService;

    @PostMapping
    public ResponseEntity<?> createProject(
            @RequestHeader("Authorization") String token,
            @RequestBody Project project
    ){
        if(token.isBlank()){
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }

        Project _project = projectService.save(project);

        return ResponseEntity.status(HttpStatus.CREATED).body(_project);
    }

    @GetMapping
    public ResponseEntity<?> findAll(
            @RequestHeader("Authorization") String token
    ){
        if(token.isBlank()){
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }

        List<Project> _projects = projectService.findAll();

        return ResponseEntity.status(HttpStatus.OK).body(_projects);
    }

    @GetMapping("/{id}")
    public ResponseEntity<?> findById(
            @PathVariable Integer id,
            @RequestHeader("Authorization") String token
    ){
        if(token.isBlank()){
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }

        Optional<Project> projectOptional = projectService.findById(id);

        if (projectOptional.isPresent()) {
            return ResponseEntity.ok(projectOptional.get());
        } else {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body("Project not found with id: " + id);
        }
    }

    @GetMapping("/title={title}")
    public ResponseEntity<?> findByTitle(
            @PathVariable String title,
            @RequestHeader("Authorization") String token
    ){
        if(token.isBlank()){
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }

        List<Project> projects = projectService.findProjectByTitle(title);
        return ResponseEntity.status(HttpStatus.OK).body(projects);
    }

}
