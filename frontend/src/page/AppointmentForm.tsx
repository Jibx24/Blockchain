import React from "react";
import { Form, Input, DatePicker, TimePicker, Button, Select, message } from "antd";
import axios from "axios";

const { Option } = Select;


const AppointmentForm: React.FC = () => {

  return (
    <div style={{ width: "100%", height: "100vh", display: "flex", justifyContent: "center", alignItems: "center" }}>
      <Form  layout="vertical" style={{ width: "80%", maxWidth: "600px", margin: "0 auto" }}>
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
          <Button type="primary" htmlType="submit" style={{ width: "100%" }}>
            Book Appointment
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default AppointmentForm;
