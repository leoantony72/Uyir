import React, { useState } from 'react';
import { UserHeader } from '../components/UserHeader';
import { ReportCard } from '../components/ReportCard';
import styles from './ReportsPage.module.css';

export const ReportsPage = () => {
  const [reports] = useState([
    {
      type: 'pothole',
      location: 'Thrissur',
      longitude: '76.21',
      latitude: '10.52',
      date: '2025-02-11',
      status: 'Fixed',
      imageUrl: '/images/report1.jpg'
    },
    {
      type: 'pothole',
      location: 'Thrissur',
      longitude: '76.21',
      latitude: '10.52',
      date: '2025-02-11',
      status: 'Pending',
      imageUrl: '/images/report2.jpg'
    },
    {
      type: 'pothole',
      location: 'Thrissur',
      longitude: '76.21',
      latitude: '10.52',
      date: '2025-02-11',
      status: 'Fixed',
      imageUrl: '/images/report3.jpg'
    }
  ]);

  const handleNewReport = () => {
    console.log('New report button clicked');
  };

  return (
    <main className={styles.container}>
      <UserHeader
        userName="Leo Antony"
        points={1000}
        avatarUrl="/images/profile-avatar.jpg"
        onNewReport={handleNewReport}
      />
      <section 
        className={styles.reportsList}
        aria-label="Road issue reports"
      >
        {reports.map((report, index) => (
          <ReportCard
            key={`report-${index}`}
            {...report}
          />
        ))}
      </section>
    </main>
  );
};