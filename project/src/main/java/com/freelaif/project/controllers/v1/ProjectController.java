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

    @GetMapping("/devId={devId}")
    public ResponseEntity<?> findByDevId(
            @PathVariable int devId,
            @RequestHeader("Authorization") String token
    ){
        if(token.isBlank()){
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }

        List<Project> projects = projectService.findByDeveloperId(devId);
        return ResponseEntity.status(HttpStatus.OK).body(projects);
    }

    @GetMapping("/clientId={clientId}")
    public ResponseEntity<?> findByclientId(
            @PathVariable int clientId,
            @RequestHeader("Authorization") String token
    ){
        if(token.isBlank()){
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }

        List<Project> projects = projectService.findByClientId(clientId);
        return ResponseEntity.status(HttpStatus.OK).body(projects);
    }

    @PutMapping("/associate")
    public ResponseEntity<?> associateDevToProject(
            @RequestBody int devId,
            @RequestBody int projectId,
            @RequestHeader("Authorization") String token){
        if(token.isBlank()){
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }
        Optional<Project> _project = projectService.associateDeveloper(projectId, devId);
        if(_project.isPresent()){
            return ResponseEntity.status(HttpStatus.OK).body(_project);
        }else{
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body("Fail to create association.");
        }
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<?> deleteProject(
            @PathVariable int id,
            @RequestHeader("Authorization") String token){
        if(token.isBlank()){
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }
        boolean _project = projectService.removeProject(id);
        if(_project){
            return ResponseEntity.status(HttpStatus.OK).body("Project removed.");
        }else{
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body("Fail to delete project.");
        }
    }

    @DeleteMapping("/developerId={developerId}&projectId={projectId}")
    public ResponseEntity<?> deleteAssociatedDeveloperFrom(
            @PathVariable int projectId,
            @RequestHeader("Authorization") String token){
        if(token.isBlank()){
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Missing token.");
        }
        boolean _project = projectService.removeDevFromProject(projectId);
        if(_project){
            return ResponseEntity.status(HttpStatus.OK).body("Dev removed from project.");
        }else{
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body("Fail to delete developer.");
        }
    }
}
