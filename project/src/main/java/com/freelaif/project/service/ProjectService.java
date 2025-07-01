package com.freelaif.project.service;

import com.freelaif.project.models.Project;
import com.freelaif.project.repository.ProjectRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class ProjectService {
    @Autowired
    private ProjectRepository projectRepository;

    @Autowired
    private SequenceGeneratorService sequenceGeneratorService;

    public Project save(Project project){
        project.setId(sequenceGeneratorService.getNextSequence("projects"));
        return projectRepository.save(project);
    }

    public Optional<Project> findById(int id){
        return projectRepository.findById(id);
    }

    public List<Project> findAll(){
        return projectRepository.findAll();
    }

    public List<Project> findProjectByTitle(String title){
        return projectRepository.findByTitle(title);
    }

    public List<Project> findByDeveloperId(int developerId){
        return projectRepository.findByDeveloperId(developerId);
    }

    public List<Project> findByClientId(int clientId){
        return projectRepository.findByClientId(clientId);
    }

    public Optional<Project> associateDeveloper(int projectId, int developerId) {
        Optional<Project> projectOptional = projectRepository.findById(projectId);
        if (projectOptional.isPresent()) {
            Project project = projectOptional.get();
            project.setDeveloperId(developerId);
            projectRepository.save(project);
            return Optional.of(project);
        }
        return Optional.empty();
    }

    public boolean removeProject(int projectId) {
        Optional<Project> projectOptional = projectRepository.findById(projectId);
        if (projectOptional.isPresent()) {
            Project project = projectOptional.get();
            projectRepository.delete(project);
            return true;
        }
        return false;
    }

    public boolean removeDevFromProject(int projectId) {
        Optional<Project> projectOptional = projectRepository.findById(projectId);
        if (projectOptional.isPresent()) {
            Project project = projectOptional.get();
            project.setDeveloperId(null);
            projectRepository.save(project);
            return true;
        }
        return false;
    }
}
