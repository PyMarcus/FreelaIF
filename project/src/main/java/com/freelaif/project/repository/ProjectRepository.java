package com.freelaif.project.repository;

import com.freelaif.project.models.Project;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface ProjectRepository extends MongoRepository<Project, String> {
    public Optional<Project> findById(Integer id);
    public List<Project> findByTitle(String title);
    public List<Project> findByDeveloperId(int developerId);
    public List<Project> findByClientId(int clientId);
}
