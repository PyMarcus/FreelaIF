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
    private ProjectService projectService;

    @PostMapping
    public ResponseEntity<?> createProject(
            @RequestHeader("Authorization") String token,
            @RequestBody Project project) {
        if (token.isBlank()) {
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }

        Project createdProject = projectService.save(project);
        return ResponseEntity.status(HttpStatus.CREATED).body(createdProject);
    }

    @GetMapping
    public ResponseEntity<?> findAll(
            @RequestHeader("Authorization") String token) {
        if (token.isBlank()) {
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }

        List<Project> projects = projectService.findAll();
        return ResponseEntity.ok(projects);
    }

    @GetMapping("/by-id/{id}")
    public ResponseEntity<?> findById(
            @PathVariable Integer id,
            @RequestHeader("Authorization") String token) {
        if (token.isBlank()) {
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }

        Optional<Project> projectOptional = projectService.findById(id);
        return projectOptional
                .<ResponseEntity<?>>map(ResponseEntity::ok)
                .orElseGet(() -> ResponseEntity.status(HttpStatus.NOT_FOUND).body("Project not found with id: " + id));
    }

    @GetMapping("/search")
    public ResponseEntity<?> searchProjects(
            @RequestParam(required = false) String title,
            @RequestParam(required = false) Integer devId,
            @RequestParam(required = false) Integer clientId,
            @RequestHeader("Authorization") String token) {
        if (token.isBlank()) {
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }

        if (title != null) {
            List<Project> projects = projectService.findProjectByTitle(title);
            return ResponseEntity.ok(projects);
        } else if (devId != null) {
            List<Project> projects = projectService.findByDeveloperId(devId);
            return ResponseEntity.ok(projects);
        } else if (clientId != null) {
            List<Project> projects = projectService.findByClientId(clientId);
            return ResponseEntity.ok(projects);
        } else {
            return ResponseEntity.badRequest().body("Please provide a search parameter: title, devId, or clientId.");
        }
    }

    @PutMapping("/{id}/developer/{devId}")
    public ResponseEntity<?> associateDevToProject(
            @PathVariable int id,
            @PathVariable int devId,
            @RequestHeader("Authorization") String token) {
        if (token.isBlank()) {
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }

        Optional<Project> updatedProject = projectService.associateDeveloper(id, devId);
        return updatedProject.<ResponseEntity<?>>map(ResponseEntity::ok)
                .orElseGet(() -> ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body("Failed to associate developer."));
    }

    @DeleteMapping("/{id}/developer")
    public ResponseEntity<?> removeDeveloperFromProject(
            @PathVariable int id,
            @RequestHeader("Authorization") String token) {
        if (token.isBlank()) {
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }

        boolean success = projectService.removeDevFromProject(id);
        if (success) {
            return ResponseEntity.ok("Developer removed from project.");
        } else {
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body("Failed to remove developer.");
        }
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<?> deleteProject(
            @PathVariable int id,
            @RequestHeader("Authorization") String token) {
        if (token.isBlank()) {
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }

        boolean success = projectService.removeProject(id);
        if (success) {
            return ResponseEntity.ok("Project removed.");
        } else {
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body("Failed to delete project.");
        }
    }
}
