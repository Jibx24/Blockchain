import React, { useState } from 'react';
import { Form, Select, DatePicker, TimePicker, Input, Button, Card, Row, Col } from 'antd';
import type { FormInstance } from 'antd/es/form';
import moment from 'moment';
const { Option } = Select;
const { TextArea } = Input;

interface DepartmentRoom {
  [key: string]: string[];
}

const AppointmentForm: React.FC = () => {
  const [form] = Form.useForm();
  const [selectedDept, setSelectedDept] = useState<string>('');

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

  // Disallow past dates
  const disabledDate = (current: moment.Moment) => {
    return current && current < moment().startOf('day');
  };

  return (
    <div style={{ padding: '24px' }}>
      <Card title="Medical Appointment Form" bordered={false}>
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
      </Card>
    </div>
  );
};

export default AppointmentForm;