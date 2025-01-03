import React from "react";
import { Table } from "antd";
import { Appointment } from "../types/Appointment";

interface Props {
  appointments: Appointment[];
}

const AppointmentTable: React.FC<Props> = ({ appointments }) => {
  const columns = [
    {
      title: "Patient ID",
      dataIndex: "data",
      key: "patient_id",
      render: (data: Appointment["data"]) => data.patient_id,
    },
    {
      title: "Doctor ID",
      dataIndex: "data",
      key: "doctor_id",
      render: (data: Appointment["data"]) => data.doctor_id,
    },
    {
      title: "Date",
      dataIndex: "data",
      key: "date",
      render: (data: Appointment["data"]) => data.date,
    },
    {
      title: "Time",
      dataIndex: "data",
      key: "time",
      render: (data: Appointment["data"]) => data.time,
    },
    {
      title: "Reason",
      dataIndex: "data",
      key: "reason",
      render: (data: Appointment["data"]) => data.reason || "-",
    },
  ];

  return <Table dataSource={appointments} columns={columns} rowKey="index" />;
};

export default AppointmentTable;
