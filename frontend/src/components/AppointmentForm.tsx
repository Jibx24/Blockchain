import React from "react";
import { Form, Input, DatePicker, TimePicker, Button, Select, message } from "antd";
import axios from "axios";

const { Option } = Select;

interface Props {
  onAppointmentAdded: () => void;
}

const AppointmentForm: React.FC<Props> = ({ onAppointmentAdded }) => {
  const [loading, setLoading] = React.useState(false);

  const onFinish = async (values: any) => {
    const { date, time, doctor_id, reason } = values;
    const appointment = {
      patient_id: "P123", // สมมติว่าดึงมาจากระบบ Authentication
      doctor_id,
      date: date.format("YYYY-MM-DD"),
      time: time.format("HH:mm"),
      reason,
    };

    setLoading(true);
    try {
      await axios.post("http://localhost:8080/appointment", appointment);
      message.success("Appointment booked successfully!");
      onAppointmentAdded();
    } catch (error) {
      message.error("Failed to book appointment");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ width: "100%", height: "100vh", display: "flex", justifyContent: "center", alignItems: "center" }}>
      <Form onFinish={onFinish} layout="vertical" style={{ width: "80%", maxWidth: "600px", margin: "0 auto" }}>
        <Form.Item label="Doctor" name="doctor_id" rules={[{ required: true, message: "Please select a doctor!" }]}>
          <Select placeholder="Select a doctor">
            <Option value="D001">Dr. Smith (Cardiology)</Option>
            <Option value="D002">Dr. Brown (Pediatrics)</Option>
            <Option value="D003">Dr. Green (Orthopedics)</Option>
          </Select>
        </Form.Item>
        <Form.Item label="Date" name="date" rules={[{ required: true, message: "Please select a date!" }]}>
          <DatePicker style={{ width: "100%" }} />
        </Form.Item>
        <Form.Item label="Time" name="time" rules={[{ required: true, message: "Please select a time!" }]}>
          <TimePicker style={{ width: "100%" }} format="HH:mm" />
        </Form.Item>
        <Form.Item label="Reason for Appointment" name="reason">
          <Input.TextArea rows={4} />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" loading={loading} style={{ width: "100%" }}>
            Book Appointment
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default AppointmentForm;
