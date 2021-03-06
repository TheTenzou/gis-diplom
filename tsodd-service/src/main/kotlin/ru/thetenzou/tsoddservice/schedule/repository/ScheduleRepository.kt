package ru.thetenzou.tsoddservice.schedule.repository

import org.springframework.data.domain.Page
import org.springframework.data.domain.Pageable
import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.data.rest.core.annotation.RepositoryRestResource
import ru.thetenzou.tsoddservice.schedule.model.Schedule

@RepositoryRestResource
interface ScheduleRepository : JpaRepository<Schedule, Long> {

    fun getAllBy(pageable: Pageable): Page<Schedule>
}