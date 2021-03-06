import React from "react";
import { Badge } from "bootstrap-4-react";

import "../css/table.css";

function ScheduleTable({ crew, schedule }) {
  let col = [<col />];
  crew?.forEach((crew) => {
    col.push(<col width="20px" />);
  });
  let firstRow = [<th></th>];

  crew?.forEach((crew) => {
    firstRow.push(<th style={{ textAlign: "center" }}>{crew.name}</th>);
  });

  let table = [
    <thead>
      <tr>{firstRow}</tr>
    </thead>,
  ];

  let tableBody = [];

  if (schedule) {
    const days = [...schedule.keys()];

    days.forEach((day) => {
      const crewsTasks = schedule.get(day);

      let dayTaskList = [<th style={{ minWidth: "88px" }}>{day}</th>];
      crew?.forEach((crew) => {
        let crewTasks = [];
        if (crewsTasks.has(crew.id)) {
          crewsTasks.get(crew.id).forEach((task) => {
            crewTasks.push(
              <div
                style={{ display: "flex", flexDirection: "row", justifyContent: "space-between", alignItems: "center", }}>
                <div>
                  Задача: {task.taskType.name}
                  <br />
                  ТСОДД: {task.tsodd.typeName} Номер: {task.tsodd.id}
                  <br />
                  Местоположение: {task.tsodd.positionDescription}
                </div>

                <div style={{ display: "flex", flexDirection: "column", justifyContent: "center", alignItems: "center" }}>
                  <Badge info className="btn btn-outline-light" style={{ marginBottom: ".2rem" }}>&#127760;</Badge>
                  <Badge primary className="btn btn-outline-light" style={{ marginBottom: ".2rem" }}>&#9998;</Badge>
                  <Badge danger className="btn btn-outline-light">&#10008;</Badge>
                </div>
              </div>
            );
          });
        }
        dayTaskList.push(<td>{crewTasks}</td>);
      });
      tableBody.push(<tr>{dayTaskList}</tr>);
    });
    table.push(<tbody>{tableBody}</tbody>);
  }

  return (
    <>
      <table className="styled-table">{table}</table>
    </>
  );
}

export default ScheduleTable;
