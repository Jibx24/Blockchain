import React from "react";
import { Table } from "antd";


const AppointmentTable: React.FC = () => {
  const columns = [
    {
      title: "Patient ID",
      dataIndex: "data",
      key: "patient_id",
    },
    {
      title: "Doctor ID",
      dataIndex: "data",
      key: "doctor_id",
    },
    {
      title: "Date",
      dataIndex: "data",
      key: "date",
    },
    {
      title: "Time",
      dataIndex: "data",
      key: "time",
    },
    {
      title: "Reason",
      dataIndex: "data",
      key: "reason",
    },
  ];

  return <Table  columns={columns} rowKey="index" />;
};

export default AppointmentTable;
