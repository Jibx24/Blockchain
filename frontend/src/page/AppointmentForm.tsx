import React, { useState } from 'react';
import { Form, Select, DatePicker, TimePicker, Input, Button, Card, Row, Col, Table } from 'antd';
import moment from 'moment';

const { Option } = Select;
const { TextArea } = Input;

interface DepartmentRoom {
  [key: string]: string[];
}

const AppointmentForm: React.FC = () => {
  const [form] = Form.useForm();
  const [selectedDept, setSelectedDept] = useState<string>('');

  // ข้อมูลสำหรับ Department และ Room
  const departments = [
    { id: 'D1', name: 'General Medicine' },
    { id: 'D2', name: 'General Surgery' },
    { id: 'D3', name: 'Obstetrics & Gynecology' },
    { id: 'D4', name: 'Pediatrics' },
    { id: 'D5', name: 'ENT' },
    { id: 'D6', name: 'Eye Clinic' },
    { id: 'D7', name: 'Orthopedics' },
    { id: 'D8', name: 'Dental Clinic' }
  ];

  const rooms: DepartmentRoom = {
    'D1': ['201', '202', '203'],
    'D2': ['301', '302', '303'],
    'D3': ['401', '402', '403'],
    'D4': ['501', '502', '503'],
    'D5': ['601', '602'],
    'D6': ['701', '702'],
    'D7': ['801', '802'],
    'D8': ['901', '902']
  };

  const onFinish = (values: any) => {
    console.log('Form values:', {
      ...values,
      date: values.date?.format('YYYY-MM-DD'),
      time: values.time?.format('HH:mm')
    });
  };

  const handleDepartmentChange = (value: string) => {
    setSelectedDept(value);
    form.setFieldsValue({ room: undefined });
  };

  // กำหนดวันไม่ให้เลือกวันในอดีต
  const disabledDate = (current: moment.Moment) => {
    return current && current < moment().startOf('day');
  };

  // ข้อมูลสำหรับตาราง
  const columns = [
    {
      title: "Patient ID",
      dataIndex: "patient_id",
      key: "patient_id",
    },
    {
      title: "Doctor ID",
      dataIndex: "doctor_id",
      key: "doctor_id",
    },
    {
      title: "Date",
      dataIndex: "date",
      key: "date",
    },
    {
      title: "Time",
      dataIndex: "time",
      key: "time",
    },
    {
      title: "Reason",
      dataIndex: "reason",
      key: "reason",
    },
  ];

  // ข้อมูลสำหรับแสดงในตาราง
  const data = [
    {
      patient_id: 'P001',
      doctor_id: 'D001',
      date: '2025-01-01',
      time: '10:00',
      reason: 'Routine Checkup',
    },
    {
      patient_id: 'P002',
      doctor_id: 'D002',
      date: '2025-01-02',
      time: '11:00',
      reason: 'Dental Consultation',
    },
  ];

  return (
    <div style={{ padding: '24px' }}>
      <Card title="Medical Appointment Form" bordered={false}>
        <Row gutter={16}>
          <Col xs={24} sm={12}>
            {/* ฟอร์มสำหรับเลือกแผนกและห้อง */}
            <Form
              form={form}
              layout="vertical"
              onFinish={onFinish}
              autoComplete="off"
            >
              <Row gutter={16}>
                <Col xs={24} sm={12}>
                  <Form.Item
                    name="department"
                    label="Department"
                    rules={[{ required: true, message: 'Please select a department' }]}
                  >
                    <Select
                      placeholder="Select department"
                      onChange={handleDepartmentChange}
                    >
                      {departments.map(dept => (
                        <Option key={dept.id} value={dept.id}>{dept.name}</Option>
                      ))}
                    </Select>
                  </Form.Item>
                </Col>

                <Col xs={24} sm={12}>
                  <Form.Item
                    name="room"
                    label="Examination Room"
                    rules={[{ required: true, message: 'Please select a room' }]}
                  >
                    <Select
                      placeholder="Select examination room"
                      disabled={!selectedDept}
                    >
                      {selectedDept && rooms[selectedDept]?.map(room => (
                        <Option key={room} value={room}>Room {room}</Option>
                      ))}
                    </Select>
                  </Form.Item>
                </Col>
              </Row>

              <Row gutter={16}>
                <Col xs={24} sm={12}>
                  <Form.Item
                    name="date"
                    label="Date"
                    rules={[{ required: true, message: 'Please select a date' }]}
                  >
                    <DatePicker 
                      style={{ width: '100%' }}
                      format="DD/MM/YYYY"
                      //disabled={disabledDate}
                    />
                  </Form.Item>
                </Col>

                <Col xs={24} sm={12}>
                  <Form.Item
                    name="time"
                    label="Time"
                    rules={[{ required: true, message: 'Please select a time' }]}
                  >
                    <TimePicker 
                      style={{ width: '100%' }}
                      format="HH:mm"
                      minuteStep={15}
                      showNow={false}
                    />
                  </Form.Item>
                </Col>
              </Row>

              <Form.Item
                name="symptoms"
                label="Initial Symptoms"
                rules={[{ required: true, message: 'Please describe your symptoms' }]}
              >
                <TextArea 
                  rows={4} 
                  placeholder="Please describe your initial symptoms"
                />
              </Form.Item>

              <Form.Item>
                <Button type="primary" htmlType="submit" block>
                  Confirm Appointment
                </Button>
              </Form.Item>
            </Form>
          </Col>

          {/* ตารางแสดงข้อมูลการนัดหมาย */}
          <Col xs={24} sm={12}>
            <Card title="Appointment Table" bordered={true}>
              <Table
                columns={columns}
                dataSource={data}
                rowKey="patient_id"
              />
            </Card>
          </Col>
        </Row>
      </Card>
    </div>
  );
};

export default AppointmentForm;
