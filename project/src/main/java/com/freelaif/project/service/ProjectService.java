package com.freelaif.project.service;

import com.freelaif.project.models.Project;
import com.freelaif.project.repository.ProjectRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ProjectService {
    @Autowired
    private ProjectRepository projectRepository;

    public Project save(Project project){
        return projectRepository.save(project);
    }

    public List<Project> findAll(){
        return projectRepository.findAll();
    }

    public List<Project> findProjectByTitle(String title){
        return projectRepository.findByTitle(title);
    }
}
